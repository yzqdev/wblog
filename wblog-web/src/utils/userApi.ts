import http from "@/utils/http";
import qs from 'qs';
export function getProfileApi(){
    return http.get('/admin/profile')
}
export function updateProfileApi(data){
    return http.postForm('/admin/profile',data)
}
