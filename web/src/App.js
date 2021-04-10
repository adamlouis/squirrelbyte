import React, { useCallback, useState, useEffect } from "react";
import _ from "lodash";
import styled from "styled-components";
import { DocumentList } from "./DocumentList";
import { subscribeKeyDown } from "./KeyPublisher";
import Util from "./Util";
import { Loader } from "./Loader";
import { Ace } from "./Ace";
import { Info } from "./Info";

const TABS = {
  result: { name: "result" },
  query: { name: "query" },
};

const Row = styled.div`
  display: flex;
  align-items: center;
  justify-content: center;
`;

const RowRight = styled(Row)`
  justify-content: flex-end;
`;

const RowLeft = styled(Row)`
  justify-content: flex-start;
`;

const TabRow = styled(RowLeft)`
  border-bottom: solid black 1px;
`;

const Tab = styled.div`
  cursor: pointer;
  padding: 10px;
  border: solid black 1px;
  margin-right: -1px;
  margin-bottom: -1px;
  background-color: ${(p) => (p.selected ? "#ccc" : "#fff")};
  :hover {
    background-color: #ccc;
  }
`;

const Banner = styled.div`
  background-color: ${(p) => p.backgroundColor};
  padding: 3px 8px;
  margin-bottom: 10px;
`;

const SubmitButton = styled.input`
  padding: 5px 10px;
  cursor: pointer;
  background-color: #ddd;
  margin-top: 10px;
`;

const QueryView = styled.pre`
  background-color: #ddd;
  padding: 20px;
  border-radius: 3px;
`;

// hand pick for demo
const defaultQuery = {
  select: [],
  where: {},
  group_by: [],
  order_by: [],
  limit: 1000,
};

const getEmptyResult = () => ({
  query: "",
  documents: undefined,
  paths: undefined,
  insights: undefined,
  error: "",
});

const getQueryFromURL = () => {
  try {
    const q = Util.getUrlParameter("q");
    if (!q) {
      Util.clearURLParameters();
      return;
    }

    const j = JSON.parse(q);

    if (!_.isObject(j)) {
      Util.clearURLParameters();
      return;
    }

    return JSON.stringify(j, undefined, 2);
  } catch (e) {
    console.warn(e);
    Util.clearURLParameters();
  }
};

const queryFromUrl = getQueryFromURL();

function App() {
  const [query, setQuery] = useState(
    queryFromUrl || JSON.stringify(defaultQuery, undefined, 2)
  );
  const [result, setResult] = useState(getEmptyResult());
  const [selectedTab, setSelectedTab] = useState("result");
  const [loading, setLoading] = useState(false);

  const onClickTab = (t) => setSelectedTab(t);
  const onSubmitForm = async (e) => {
    e.preventDefault();
    submitForm();
  };

  const submitForm = useCallback(async () => {
    setLoading(true);
    setResult(getEmptyResult());

    const submittedQuery = query;
    try {
      const start = performance.now();
      const res = await fetch("/api/documents:search", {
        method: "POST",
        body: submittedQuery,
        json: true,
      });
      const elapsed = performance.now() - start;

      const j = await res.json();

      setResult({
        query: submittedQuery,
        documents: j.result,
        paths: Util.getAllPaths(j.result),
        insights: j.insights,
        error: j.message,
        elapsed,
      });

      try {
        Util.setUrlParameter("q", JSON.stringify(JSON.parse(submittedQuery)));
      } catch (e) {
        console.warn(e);
      }
    } catch (e) {
      setResult({
        query: submittedQuery,
        paths: [],
        documents: [],
        insights: {},
        error: `${e}`,
      });
    }
    setLoading(false);
  }, [setLoading, setResult, query]);

  useEffect(() => {
    if (queryFromUrl) {
      submitForm();
    }
    // eslint-disable-next-line react-hooks/exhaustive-deps
  }, []);

  useEffect(() => {
    const unsubscribeCmdB = subscribeKeyDown("KeyB", true, () => {
      setQuery(Util.safePretty(query));
    });
    const unsubscribeCmdEnter = subscribeKeyDown("Enter", true, () => {
      if (loading) {
        return;
      }
      submitForm();
    });
    return () => {
      unsubscribeCmdB();
      unsubscribeCmdEnter();
    };
  }, [query, submitForm, loading]);

  let banner;
  if (result.error) {
    banner = (
      <Banner backgroundColor={Util.Colors.Yellow}>{result.error}</Banner>
    );
  } else {
    if (result.query && result.insights) {
      banner = (
        <Banner backgroundColor={Util.Colors.Green}>{`${_.size(
          _.get(result, "documents")
        )} records - ${Math.round(
          _.get(result, "insights.duration_ms")
        )}ms service time - ${Math.round(
          result.elapsed
        )}ms round trip time`}</Banner>
      );
    }
  }

  return (
    <div
      style={{
        margin: "0px 30px 30px 30px",
      }}
    >
      <div>
        <div
          style={{
            display: "flex",
            alignItems: "center",
            justifyContent: "flex-start",
            marginTop: "15px",
          }}
        >
          <div
            style={{
              fontSize: "20px",
              fontWeight: "bold",
            }}
          >
            squirrel byte
          </div>
          <span
            style={{
              backgroundColor: "gold",
              fontSize: "14px",
              padding: "5px",
              marginLeft: "8px",
              fontWeight: "bold",
            }}
          >
            proof of concept
          </span>
        </div>
        <Info />
      </div>
      <form
        style={{
          display: "flex",
          flexDirection: "column",
          alignItems: "center",
        }}
        onSubmit={onSubmitForm}
      >
        <div
          style={{
            width: "100%",
            margin: "10px 0px",
          }}
        >
          <Ace
            initialValue={query}
            height={"400px"}
            width={"100%"}
            onChange={setQuery}
          />
          <RowRight>
            <SubmitButton type="submit" value="submit" />
          </RowRight>
        </div>
      </form>
      <div style={{ minHeight: "750px" }}>
        {loading && (
          <div
            style={{
              width: "100%",
              height: "100%",
              display: "flex",
              justifyContent: "center",
            }}
          >
            <Loader size={"25px"} borderSize={"5px"} />
          </div>
        )}
        {banner}
        {result.documents && (
          <div>
            <TabRow>
              {_.map(TABS, (v, k) => (
                <Tab
                  onClick={() => onClickTab(k)}
                  key={k}
                  selected={k === selectedTab}
                >
                  {v.name}
                </Tab>
              ))}
            </TabRow>
            {selectedTab === "query" && (
              <div style={{ display: "flex" }}>
                <QueryView>{Util.safePretty(result.query)}</QueryView>
              </div>
            )}
            {selectedTab === "result" && (
              <div
                style={{
                  margin: "10px 0px",
                }}
              >
                <DocumentList
                  paths={result.paths}
                  documents={result.documents}
                />
              </div>
            )}
          </div>
        )}
      </div>
    </div>
  );
}

export default App;
