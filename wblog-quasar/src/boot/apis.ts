import {api} from './axios';

export const getIndex = () => {
    return api.get("/home/index");
};

export const getPosts = () => {
  return api.get("/home/posts");
};
export const getPostById = (id:string) => {
  return api.get(`/home/post/${id}`);
};
