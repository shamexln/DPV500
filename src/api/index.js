import requests  from "./request";

export const captureNewImage = (params)=> {
    return requests({url:'images', method: 'post',data:params});
}