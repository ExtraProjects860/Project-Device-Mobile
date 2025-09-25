import { ApiDefault, instanceNoticesApi } from "./axios.js";

class NoticesRequests extends ApiDefault {
  constructor(axiosInstance) {
    super(axiosInstance);
  }

  GetNoticesRequest() {
    return;
  }
}

export default new NoticesRequests(instanceNoticesApi);
