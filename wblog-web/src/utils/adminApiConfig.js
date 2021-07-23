import axios from "@/utils/axios";

export const categoryCreate = () => {
    // 新增分类
    return axios.post("/admin/categories/create");
};
export const categoryUpdate = () => {
    // 编辑分类
    return axios.post("/admin/categories/update");
};
export const getAdminCategories = () => {
    return axios.post("/admin/categories");
};
export const categoryStatus = () => {
    return axios.post("/admin/category/status/update");
};
export const getAdminArticles = () => {
    return axios.post("/admin/articles");
};
export const updateArticleStatus = () => {
    // 更新文章状态
    return axios.post("/admin/articles/status/update");
};
export const getComments = () => {
    return axios.post("/admin/comments");
};
export const updateCommentStatus = () => {
    // 更新评论状态
    return axios.post("/admin/comments/update/status/:id");
};
export const getAdminUserList = () => {
    return axios.post("/admin/users");
};
export const crawl = () => {
    // 抓取文章
    return axios.post("/admin/crawl");
};
export const customCrawl = () => {
    // 自定义抓取
    return axios.post("/admin/customcrawl");
};
export const getCrawlAccount = () => {
    return axios.post("/admin/crawl/account"); // 获取爬虫账号
};
export const createCrawlAccount = () => {
    return axios.post("/admin/crawl/account"); // 获取爬虫账号
};
export const pushToBaidu = () => {
    return axios.post("/admin/pushBaiduLink"); // 链接提交到百度
};
