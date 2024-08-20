axios.defaults.baseURL = 'http://192.168.0.20:8080';


// 创建一个 Axios 实例
const instance = axios.create({
    baseURL: axios.defaults.baseURL, // 接口的基础 URL
    timeout: 5000, // 请求超时时间
    headers: {
        'Content-Type': 'application/json', // 请求头信息
    },
});

// 添加请求拦截器
instance.interceptors.request.use(
    (config) => {
        // 在发送请求之前可以做一些操作，如添加 token 等
        return config;
    },
    (error) => {
        return Promise.reject(error);
    }
);

// 添加响应拦截器
instance.interceptors.response.use(
    (response) => {
        // 在处理响应数据之前可以做一些操作，如统一处理错误码等
        return response;
    },
    (error) => {
        // 统一处理异常
        if (error.response) {
            switch (error.response.status) {
                case 401:
                    // 未授权，跳转到登录页面
                    break;
                case 404:
                    break;
                case 500:
                    // 服务器内部错误
                    break;
                default:
                    // 其他错误
                    break;
            }
        }
        return Promise.reject(error);
    }
);

function get(url, params) {
    return instance.get(url, params);
}

function post(url, data) {
    return instance.post(url, data);
}

function download(url = '', fileName = '未知文件') {
    const a = document.createElement('a');
    a.style.display = 'none';
    a.setAttribute('target', '_blank');
    fileName && a.setAttribute('download', fileName);
    a.href = axios.defaults.baseURL+"/download?fileName="+url;
    document.body.appendChild(a);
    a.click();
    document.body.removeChild(a);
}
