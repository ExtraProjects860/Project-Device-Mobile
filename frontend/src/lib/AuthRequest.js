import { ApiDefault, instanceMainApi } from "./axios.js";

class AuthRequests extends ApiDefault {
  constructor(axiosInstance) {
    super(axiosInstance);
  }

  RequestToken() {
    return;
  }

  ResetPasswordRequest() {
    return;
  }

  LoginRequest() {
    return;
  }

  RefreshTokenRequest() {
    return;
  }

  LogoutRequest() {
    return;
  }
}

export default new AuthRequests(instanceMainApi);
