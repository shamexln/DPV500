import axios from "axios";

const requests = axios.create ({
    baseURL:"/api",
    timeout:5000,
})

requests.interceptors.request.use((config)=>{
    return config;
})


// 响应拦截器
requests.interceptors.response.use((res)=>{
    return res.data;
},
(error)=>{
    console.log(error);
    alert("服务器响应数据失败");
})
export default requests;