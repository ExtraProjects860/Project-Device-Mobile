class Product {
    #name;
    #description;
    #value;
    #discount;
    #quantity;
    #isAvailable;
    #isPromotionAvailable;
    #photoUrl;

    constructor(
        name,
        description,
        value,
        discount,
        quantity,
        isAvailable,
        isPromotionAvailable,
        photoUrl,
    ) {
        this.name = name;
        this.description = description;
        this.value = value;
        this.discount = discount;
        this.quantity = quantity;
        this.isAvailable = isAvailable;
        this.isPromotionAvailable = isPromotionAvailable;
        this.photoUrl = photoUrl;
    }

    get name() {
        return this.#name;
    }
    set name(value) {
        this.#name = value;
    }

    get description() {
        return this.#description;
    }
    set description(value) {
        this.#description = value;
    }

    get value() {
        return this.#value;
    }
    set value(val) {
        this.#value = val ? parseFloat(val) : 0;
    }

    get discount() {
        return this.#discount;
    }
    set discount(val) {
        this.#discount = val ? parseFloat(val) : 0;
    }

    get quantity() {
        return this.#quantity;
    }
    set quantity(val) {
        this.#quantity = val ? parseInt(val, 10) : 0;
    }

    get isAvailable() {
        return this.#isAvailable;
    }
    set isAvailable(value) {
        this.#isAvailable = Boolean(value);
    }

    get isPromotionAvailable() {
        return this.#isPromotionAvailable;
    }
    set isPromotionAvailable(value) {
        this.#isPromotionAvailable = Boolean(value);
    }

    get photoUrl() {
        return this.#photoUrl;
    }
    set photoUrl(value) {
        this.#photoUrl = value ? value : null;
    }

    static validateField(fieldName, value) {
        let error = null;

        switch (fieldName) {
            case "name":
                if (!value) error = "Nome é obrigatório.";
                else if (value.length < 3)
                    error = "Nome deve ter pelo menos 3 caracteres.";
                break;
            case "description":
                if (!value) error = "Descrição é obrigatória.";
                break;
            case "value":
                const numValue = parseFloat(value);
                if (!value && value !== 0) error = "Valor é obrigatório.";
                else if (isNaN(numValue) || numValue <= 0)
                    error = "Valor deve ser um número positivo.";
                break;
            case "quantity":
                const numQuantity = parseInt(value, 10);
                if (value === null || value === undefined || value === "")
                    error = "Quantidade é obrigatória.";
                else if (isNaN(numQuantity) || numQuantity < 0)
                    error = "Quantidade deve ser um número não-negativo.";
                break;
            case "discount":
                const numDiscount = parseFloat(value);
                if (value && (isNaN(numDiscount) || numDiscount < 0 || numDiscount > 1))
                    error = "Desconto deve ser um número entre 0 (0%) e 1 (100%).";
                break;
        }
        return error;
    }

    static validateAll(formData) {
        const errors = {};
        const { name, description, value, quantity, discount } = formData;

        errors.name = Product.validateField("name", name);
        errors.description = Product.validateField("description", description);
        errors.value = Product.validateField("value", value);
        errors.quantity = Product.validateField("quantity", quantity);
        errors.discount = Product.validateField("discount", discount);

        return errors;
    }

    static getChangedFields(originalProduct, newFormData) {
        const updatedProductData = {};
        const {
            name,
            description,
            value,
            discount,
            quantity,
            isAvailable,
            isPromotionAvailable,
            photoUrl,
        } = newFormData;

        const valueFloat = value ? parseFloat(value) : 0;
        const discountFloat = discount ? parseFloat(discount) : 0;
        const quantityInt = quantity ? parseInt(quantity, 10) : 0;
        const boolIsAvailable = Boolean(isAvailable);
        const boolIsPromotionAvailable = Boolean(isPromotionAvailable);

        if (name !== originalProduct.name) {
            updatedProductData.name = name;
        }
        if (description !== originalProduct.description) {
            updatedProductData.description = description;
        }
        if (valueFloat !== originalProduct.value) {
            updatedProductData.value = valueFloat;
        }
        if (discountFloat !== originalProduct.discount) {
            updatedProductData.discount = discountFloat;
        }
        if (quantityInt !== originalProduct.quantity) {
            updatedProductData.quantity = quantityInt;
        }

        if (boolIsAvailable !== originalProduct.is_avaible) {
            updatedProductData.is_available = boolIsAvailable;
        }
        if (
            boolIsPromotionAvailable !== originalProduct.is_promotion_avaible
        ) {
            updatedProductData.is_promotion_available = boolIsPromotionAvailable;
        }

        if (photoUrl !== originalProduct.photo_url) {
            updatedProductData.photo_url = photoUrl || "";
        }

        return updatedProductData;
    }

    toJSON() {
        return {
            name: this.name,
            description: this.description,
            value: this.value,
            discount: this.discount,
            quantity: this.quantity,
            is_available: this.isAvailable,
            is_promotion_available: this.isPromotionAvailable,
            photo_url: this.photoUrl,
        };
    }
}

export default Product;