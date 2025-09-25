import { ApiDefault, instanceMainApi } from "./axios.js";

class WishListRequest extends ApiDefault {
  constructor(axiosInstance) {
    super(axiosInstance);
  }

  AddInWishListRequest() {
    return;
  }

  UpdateWishListRequest() {
    return;
  }

  GetItemsWishListRequest() {
    return;
  }
}

export default new WishListRequest(instanceMainApi);
