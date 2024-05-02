import Axios from "axios";
import { toast } from "sonner";

const api = Axios.create({
  baseURL: "/api",
  withCredentials: true,
});

api.interceptors.response.use(
  (response) => {
    // return response.data;
    return response;
  },
  (error) => {
    const res = error.response;

    if (res.status >= 500) {
      const message = res?.data?.message || error.message;
      toast(res?.statusText, {
        description: message,
      });
    }
    return Promise.reject(error);
  },
);

export default api;
