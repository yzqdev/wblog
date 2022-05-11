import http from "./http";
export const postApi = (msg: any) => {
  return http.postForm("/admin/posts", msg);
};
export const loginApi = (msg: any) => {
  return http.post("/auth/login", msg);
};
export const regApi = (msg: any) => {
  return http.post("/auth/reg", msg);
};
export  const getPostsApi=( ) => {
  return http.get("/admin/posts")
}
export  const delPostsApi=(id:string ) => {
  return http.delete(`/admin/posts/${id}`)
}
