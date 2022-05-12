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
export const getLinks = ( ) => {
  return api.get(`/home/links`);
};
//comment

export const postComment = (postId:string,data:any ) => {
  return api.postForm(`/home/comment/${postId}`,data);
};
