import { React, useState } from "react";
import { Queries } from "./Queries";
import styled from "styled-components";

const InfoRow = styled.div`
  width: 100%;
  display: flex;
  align-items: center;
  justify-content: space-between;
  cursor: pointer;
  :hover {
    color: #555;
  }
`;

// prettier adding weird formatting but it's fine
export function Info() {
  const [show, setShow] = useState(!!localStorage.getItem("show-description"));

  const onClickShow = (e) => {
    e.stopPropagation();
    localStorage.setItem("show-description", !show ? "true" : "");
    setShow(!show);
  };

  return (
    <div
      style={{
        backgroundColor: "#ddd",
        padding: "2px 20px",
        margin: "15px 0px 10px 0px",
        fontSize: "14px",
        position: "relative",
      }}
    >
      <InfoRow onClick={onClickShow}>
        <p>
          <strong>author's note</strong>
        </p>
        <button
          style={{
            padding: "5px 10px",
            cursor: "pointer",
          }}
          onClick={onClickShow}
        >
          {show ? "hide" : "show"}
        </button>
      </InfoRow>
      <div style={{ display: show ? "" : "none" }}>
        <p>
          <strong>why</strong>
          <br />
          <br />I like to explore & understand data from the software services I
          use - Strava, Garmin, GitHub, AWS, & some others.
          <br />
          <br />
          Some tools I like for exploring data in general are:
          <ul>
            <li>
              <a target="_blank" rel="noreferrer" href="https://datasette.io/">
                Datasette
              </a>{" "}
              - for exploring sqlite databases
            </li>
            <li>
              <a
                target="_blank"
                rel="noreferrer"
                href="https://stedolan.github.io/jq/"
              >
                jq
              </a>{" "}
              - for sifting through local json files
            </li>
            <li>
              <a
                target="_blank"
                rel="noreferrer"
                href="https://www.honeycomb.io/overview/"
              >
                Honeycomb
              </a>{" "}
              - for general observability of distributed systems ... but in this
              case, for the query UI & how it works nicely for high cardinality
              data.
            </li>

            <li>
              <a target="_blank" rel="noreferrer" href="https://grafana.com/">
                Grafana
              </a>{" "}
              /{" "}
              <a
                target="_blank"
                rel="noreferrer"
                href="https://www.elastic.co/demos"
              >
                ElasticSearch + Kibana
              </a>{" "}
              - for general dashboard building, data ingestion, etc.
            </li>
          </ul>
          For my usecase, I wanted a way to:
          <ul>
            <li>Stash my data in its "original" JSON form</li>
            <li>Explore it later & build whatever views I want</li>
            <li>Keep costs & infrastructure complexity low</li>
            <li>Self-host it / own my data</li>
          </ul>
          The web application here is a first step towards these goals -- a
          document / search server & UI, drawing inspiration from tools that I
          like. What's here is minimal, but could become more.
        </p>
        <p>
          <strong>what</strong>
          <br />
          <br />
          squirrel byte is a "proof of concept" document / search server backed
          by sqlite. JSON documents are stashed in sqlite using the{" "}
          <a
            target="_blank"
            rel="noreferrer"
            href="https://www.sqlite.org/json1.html"
          >
            sqlite json1
          </a>{" "}
          extension. It supports a query syntax similar to{" "}
          <a target="_blank" rel="noreferrer" href="https://jsonlogic.com/">
            jsonlogic
          </a>
          , which I basically use as a (restricted) AST for a SQL query. The
          server is written in golang & the web application in js/react. It uses{" "}
          <a
            target="_blank"
            rel="noreferrer"
            href="https://bvaughn.github.io/react-virtualized/"
          >
            react-virtualized
          </a>{" "}
          so rendering thousands of rows is performant. The query "UI" is a JSON
          editor, but could be nice with{" "}
          <a
            target="_blank"
            rel="noreferrer"
            href="react-awesome-query-builder "
          >
            react-awesome-query-builder
          </a>{" "}
          or Honeycomb-like or Kibana-like forms. Could be interesting to use as
          a store for{" "}
          <a target="_blank" rel="noreferrer" href="https://mitmproxy.org/">
            mitmproxy
          </a>{" "}
          or equivalent to search network traffic.{" "}
          <a target="_blank" rel="noreferrer" href="https://duckdb.org/">
            DuckDB
          </a>{" "}
          seems cool as single file OLAP db but no JSON support. This server is
          running in read-only mode - partially because the flexible query
          syntax I want to support will not work with prepared SQL statements.
          There are a bunch of cool tools like
          <a target="_blank" rel="noreferrer" href="https://litestream.io/">
            litestream
          </a>{" "}
          or <a href="https://github.com/rqlite/rqlite">rqlite</a>
          that could help productionize sqlite if I decide to do anything
          further with this project. Has someone done everything I've described
          already?
          <br />
          <br />I typically toil away on personal software projects in private
          or spend my days working on close-source applications -- I'm starting
          to share more openly, even if it's the{" "}
          <a
            target="_blank"
            rel="noreferrer"
            href="https://github.com/adamlouis/squirrelbyte"
          >
            spaghetti here on GitHub.
          </a>
          <br />
          <br />
          <strong>weird name</strong>
          <br />
          <br />
          yes but it kind of works ... "sql" is a subsequence of "squirrel" ...
          "squirrel" means "to store up for future use" & that's what we're
          doing with our / bytes ... domain name was available
        </p>
      </div>
      <div>
        <p>
          <strong>example</strong>
          <br />
          <br />
          As a demo, this applications offers ~10k recent HackerNews article
          submissions. The `body` field of each JSON document is the raw value
          from the{" "}
          <a
            target="_blank"
            rel="noreferrer"
            href="https://github.com/HackerNews/API"
          >
            HackerNews API
          </a>
          . The links below load some queries I've prepared:
          <ul>
            {Queries.map((q) => (
              <li>
                <div key={q.name}>
                  <a href={`/?q=${encodeURIComponent(JSON.stringify(q.q))}`}>
                    {q.name}
                  </a>
                </div>
              </li>
            ))}
          </ul>{" "}
        </p>
      </div>
    </div>
  );
}
