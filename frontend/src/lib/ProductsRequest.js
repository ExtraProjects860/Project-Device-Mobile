import { ApiDefault, instanceMainApi } from "./axios.js";

class ProductRequest extends ApiDefault {
  constructor(axiosInstance) {
    super(axiosInstance);
  }

  CreateProductRequest() {
    return;
  }

  UpdateProductRequest() {
    return;
  }

  GetProductsRequest() {
    return;
  }
}

export default new ProductRequest(instanceMainApi);
