import { ApiDefault, instanceMainApi } from "./axios.js";

class UserRequest extends ApiDefault {
  constructor(axiosInstance) {
    super(axiosInstance);
  }

  CreateUserRequest() {
    return;
  }

  GetInfoUserRequest() {
    return;
  }

  GetUsersRequest() {
    return;
  }

  UpdateUserRequest() {
    return;
  }
}

export default new UserRequest(instanceMainApi);
