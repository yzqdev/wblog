import http from "@/utils/http";
import qs from 'qs';

export function superIndexApi(){
    return http.get("/super/index")
}
