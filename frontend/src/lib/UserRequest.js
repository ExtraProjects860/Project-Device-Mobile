import { ApiDefault, instanceMainApi } from "./axios.js";

class UserRequest extends ApiDefault {
  constructor(axiosInstance) {
    super(axiosInstance);
  }

  async GetUsersRequest(itemsPerPage = 20, currentPage = 1) {
    const response = await this.axiosInstance.get(
      `/users?itemsPerPage=${itemsPerPage}&currentPage=${currentPage}`
    );
    return response.data;
  }

  CreateUserRequest() {
    return;
  }

  GetInfoUserRequest() {
    return;
  }

  UpdateUserRequest() {
    return;
  }
}

export default new UserRequest(instanceMainApi);
