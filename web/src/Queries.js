export const Queries = [
  {
    name: 'top stories',
    q: {
      order_by: [{ desc: { var: 'body.score' } }],
      limit: 1000,
    },
  },
  {
    name: 'top stories of 2021-03-14 or 2021-04-01',
    q: {
      where: {
        or: [
          {
            and: [
              { '>=': [{ var: 'body.time' }, 1615708800] },
              { '<': [{ var: 'body.time' }, 1615791600] },
            ],
          },
          {
            and: [
              { '>=': [{ var: 'body.time' }, 1617260400] },
              { '<': [{ var: 'body.time' }, 1617346800] },
            ],
          },
        ],
      },
      order_by: [{ desc: { var: 'body.score' } }],
      limit: 1000,
    },
  },
  {
    name: 'top `Show HN` stories',
    q: {
      where: {
        like: [
          'Show HN:%',
          {
            var: 'body.title',
          },
        ],
      },
      order_by: [
        {
          desc: {
            var: 'body.score',
          },
        },
      ],
      limit: 1000,
    },
  },
  {
    name: 'top scoring submitters',
    q: {
      select: [
        {
          as: [
            {
              var: 'body.by',
            },
            'submitter',
          ],
        },
        {
          as: [
            {
              sum: [
                {
                  var: 'body.score',
                },
              ],
            },
            'total_score',
          ],
        },
        {
          as: [
            {
              count: [1],
            },
            'story_count',
          ],
        },
      ],
      order_by: [
        {
          desc: {
            sum: [
              {
                var: 'body.score',
              },
            ],
          },
        },
      ],
      group_by: [
        {
          var: 'body.by',
        },
      ],
      limit: 100,
    },
  },
  {
    name: 'top stories with titles containing rust, golang, or sqlite',
    q: {
      where: {
        or: [
          {
            like: [
              '%rust%',
              {
                var: 'body.title',
              },
            ],
          },
          {
            like: [
              '%golang%',
              {
                var: 'body.title',
              },
            ],
          },
          {
            like: [
              '%sqlite%',
              {
                var: 'body.title',
              },
            ],
          },
        ],
      },
      order_by: [
        {
          desc: {
            var: 'body.score',
          },
        },
      ],
      limit: 100,
    },
  },

  {
    name: 'top stories containing youtube urls',
    q: {
      where: {
        or: [
          {
            like: [
              '%https://youtube.com%',
              {
                var: 'body',
              },
            ],
          },
        ],
      },
      order_by: [
        {
          desc: {
            var: 'body.score',
          },
        },
      ],
      limit: 100,
    },
  },
];
