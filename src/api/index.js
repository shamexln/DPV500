import requests  from "./request";

export const captureNewImage = (params)=> {
    return requests({url:'images', method: 'post',data:params});
}

export const captureDeleteImage = (fileName)=>requests({url:`/images/${fileName}`,method:'delete'});
