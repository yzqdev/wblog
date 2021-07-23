<template>
  <div>
    <div class="home-categoties-box">
      <a
        href="/"
        class="categoties-item"
        :class="{ 'categoties-select': !cate }"
        >全部</a
      >
      <a
        v-for="cateItem in categories"
        class="categoties-item"
        :href="'/?cate=' + cateItem.id"
        :class="{ 'categoties-select': cateItem.id == cate }"
        >{{ cateItem.name }}</a
      >
    </div>
    <div class="home-articles-box">
      <div v-for="article in articles" class="articles-cell">
        <a
          :href="'/user/' + article.user.id"
          target="_blank"
          class="user-icon-box"
          ><img :src="article.user.avatarURL" alt=""
        /></a>
        <span class="home-tip-container">
          <el-tooltip
            :content="`回复数${article.commentCount}　浏览数${article.browseCount}`"
            placement="bottom-start"
            class="home-tip-box"
          >
            <a
              :href="'/topic/' + article.id"
              target="_blank"
              class="no-underline"
            >
              <span class="articles-click-num">{{ article.commentCount }}</span>
              <span class="articles-num-split">/</span>
              <span class="articles-res-num">{{ article.browseCount }}</span>
            </a>
          </el-tooltip>
        </span>
        <span
          class="articles-categoties"
          :class="
            article.isTop
              ? 'articles-categoties-top'
              : 'articles-categoties-common'
          "
          >{{ article.isTop ? "置顶" : article.categories[0].name }}</span
        >
        <a
          :href="'/topic/' + article.id"
          target="_blank"
          class="home-articles-title"
          :title="article.name"
          >{{    entity2HTML(article.name) }}</a
        >
        <p class="articles-res-time">{{ getReplyTime(article.createdAt) }}</p>
        <a
          v-if="article.lastUser && article.lastUser.id"
          :href="'/user/' + article.lastUser.id"
          target="_blank"
          class="user-small-icon-box"
          ><img :src="article.lastUser.avatarURL" alt=""
        /></a>
      </div>

      <div v-if="articles.length > 0" style="text-align: center">
        <span
          v-if="totalVisible"
          class="ivu-page-total"
          style="margin-top: 10px; vertical-align: top"
          >共 {{ totalCount }} 条</span
        >
        <el-pagination
          class="common-page"
          :class="{ 'common-page-inline': totalVisible }"
          :current="pageNo"
          :page-size="pageSize"
          :total="totalCount"
          @on-change="onPageChange"
          :show-elevator="true"
        />
      </div>
    </div>
  </div>
</template>

<script>
import dateTool from "@/utils/date";
import htmlUtil from "@/utils/html";
import { getArticles, getCategories, getTopList } from "@/utils/apiConfig";

export default {
  name: "PageIndex",
  data() {
    return {
      query: { cate: "", pageNo: "" },
      articles: [],
      totalVisible: false,
      pageSize: 4,
      totalCount: "",
      pageNo: 1,
      user: "",
      cate: "",
      categories: [],
    };
  },
  created() {
    this.$store.commit("top10Visible", true);
    this.$store.commit("friendLinkVisible", true);
    this.$store.commit("statVisible", true);

    Promise.all([
      getCategories(),
      getArticles({
        cateId: this.query.cate || "",
        pageNo: this.query.pageNo || 1,
        noContent: "true",
      }),
      getTopList(),
    ])
      .then((data) => {
        let user = context.user;
        let cate = this.query.cate || "";
        let categories = data[0].data.categories || [];
        let articles = data[1].data.articles || [];
        let pageNo = data[1].data.pageNo;
        let totalCount = data[1].data.totalCount;
        let pageSize = data[1].data.pageSize;
        let topList = (data[2] && data[2].data.articles) || [];
        articles.map((items) => {
          items.isTop = false;
        });
        if (!this.query.pageNo || parseInt(this.query.pageNo) < 2) {
          topList.map((items) => {
            items.isTop = true;
          });
          articles = topList.concat(articles);
        }

        this.totalVisible = process.env.NODE_ENV !== "production";
        this.categories = categories;
        this.articles = articles;
        this.totalCount = totalCount;
        this.pageNo = pageNo;
        this.pageSize = pageSize;
        this.user = user;
        this.cate = cate;
      })
      .catch((err) => {
        console.log(err.message);
        this.$message.error({ message: "Not Found", statusCode: 404 });
      });
  },

  methods: {
    onPageChange(value) {
      window.location.href = `/?cate=${this.cate}&pageNo=${value}`;
    },
    getReplyTime: dateTool.getReplyTime,
    entity2HTML: htmlUtil.entity2HTML,
  },
};
</script>

<style>
@import "../assets/styles/home.css";
</style>
