<template>
  <div class="common-body-box">
    <div class="common-body-main">
      <ul class="common-body-nav">
        <li>
          <a href="/"><span>主页</span></a>
        </li>
        <li class="common-body-nav-sep"><span>/</span></li>
        <li><span class="top100-label">Top 100 积分榜</span></li>
      </ul>
      <div class="rank-container">
        <el-table
            class="rank-list"
            :rowClassName="rowClassName"
            :data="topUsers"
            :columns="columns"
        />
      </div>
    </div>
  </div>
</template>

<script>

import {getTop100} from "@/utils/apiConfig";

export default {
  name: "Rank",
  data() {
    return {
      topUsers:'',
      columns: [
        {
          title: "#",
          key: "index",
        },
        {
          title: "用户名",
          key: "name",
          render: (h, obj) => {
            return h(
                "a",
                {
                  class: "rank-line-list-link",
                  attrs: {
                    href: `/user/${obj.row.id}`,
                  },
                },
                [
                  h("img", {
                    class: "rank-line-list-img",
                    attrs: {
                      src: obj.row.avatarURL,
                    },
                  }),
                  h(
                      "span",
                      {
                        class: "rank-line-list",
                      },
                      obj.row.name
                  ),
                ]
            );
          },
        },
        {
          title: "积分",
          key: "score",
        },
        {
          title: "话题数",
          key: "articleCount",
        },
        {
          title: "回复数",
          key: "commentCount",
        },
      ],
    };
  },
  created() {

      getTop100()

        .then((data) => {
          let user = context.user;
          this.topUsers = data[0].data.users || [];
          topUsers.map((item, index) => {
            item.index = index + 1;
          });
          return {
            user: user,
            topUsers: topUsers,
          };
        })
        .catch((err) => {
          console.log(err);
          context.error({statusCode: 404, message: "Page not found"});
        });
  },
  methods: {
    rowClassName: (row, index) => {
      return index % 2 ? "" : "rank-line-active";
    },
  },
  head() {
    return {
      title: "积分榜",
    };
  },
};
</script>

<style>
@import "../assets/styles/rank/index.css";
</style>
