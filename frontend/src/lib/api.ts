import axios from "axios";

const api = axios.create({
  baseURL: "/api",
  withCredentials: true,
});

// let accessToken = '';
//
// function getAccessToken() {
//   return accessToken;
// }
//
// // // Add a request interceptor
// api.interceptors.request.use(
//   async function (config) {
//     const savedToken = getAccessToken();
//     if (savedToken !== '') {
//       config.headers.Authorization = `Bearer ${accessToken}`;
//     } else {
//       const response = await axios.post(config.baseURL + '/auth/refresh', {});
//
//       if (response.data) {
//         accessToken = response.data.accessToken;
//         config.headers.Authorization = `Bearer ${accessToken}`;
//       }
//     }
//     return config;
//   },
//   function (error) {
//     // Do something with request error
//     return Promise.reject(error);
//   },
// );
//
// // Add a response interceptor
// api.interceptors.response.use(
//   function (response) {
//     // Any status code that lie within the range of 2xx cause this function to trigger
//     // Do something with response data
//     // const accessToken = response.headers["x-access-token"];
//     // if (accessToken) {
//     //   state.setAccessToken(accessToken);
//     // }
//     return response;
//   },
//   async function (error) {
//     // Any status codes that falls outside the range of 2xx cause this function to trigger
//     // Do something with response error
//     // Check if the error is due to an expired token
//     if (error.response && error.response.status === 401) {
//       // Refresh the access token
//       try {
//         const response = await axios.post(
//           error.response.config.baseURL + '/auth/refresh',
//           {},
//         );
//         if (response.data) {
//           accessToken = response.data.accessToken;
//           error.config.headers.Authorization = `Bearer ${accessToken}`;
//         }
//         return error;
//       } catch (refreshError) {
//         return Promise.reject(refreshError);
//       }
//     }
//     return Promise.reject(error);
//   },
// );

export default api;
