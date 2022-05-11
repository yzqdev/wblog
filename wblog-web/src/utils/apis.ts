import http from "./http";
export const notice = (msg: string) => {
  return http.post("/notice", msg);
};
export const loginApi = (msg: any) => {
  return http.post("/auth/login", msg);
};
export const regApi = (msg: any) => {
  return http.post("/auth/reg", msg);
};
