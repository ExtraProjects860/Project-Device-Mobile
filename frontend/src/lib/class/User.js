class User {
  #name;
  #cpf;
  #email;
  #registerNumber;
  #passWord;
  #roleId;
  #enterpriseId;
  #photoAsset;

  constructor(
    name,
    cpf,
    email,
    registerNumber,
    roleId,
    enterpriseId,
    photoAsset,
  ) {
    this.name = name;
    this.cpf = cpf;
    this.email = email;
    this.registerNumber = registerNumber;
    this.passWord = cpf;
    this.roleId = roleId;
    this.enterpriseId = enterpriseId;
    this.photoAsset = photoAsset;
  }

  get name() {
    return this.#name;
  }
  set name(value) {
    this.#name = value;
  }

  get cpf() {
    return this.#cpf;
  }
  set cpf(value) {
    this.#cpf = value.replace(/\D/g, "");
  }

  get email() {
    return this.#email;
  }
  set email(value) {
    this.#email = value;
  }

  get registerNumber() {
    return this.#registerNumber;
  }
  set registerNumber(value) {
    this.#registerNumber = value;
  }

  get passWord() {
    return this.#passWord;
  }
  set passWord(value) {
    this.#passWord = value;
  }

  get roleId() {
    return this.#roleId;
  }
  set roleId(value) {
    this.#roleId = parseInt(value, 10);
  }

  get enterpriseId() {
    return this.#enterpriseId;
  }
  set enterpriseId(value) {
    this.#enterpriseId = value ? parseInt(value, 10) : null;
  }

  get photoAsset() {
    return this.#photoAsset;
  }
  set photoAsset(value) {
    this.#photoAsset = value ? value : null;
  }

  /**
   * Valida um único campo.
   * @param {string} fieldName
   * @param {string} value
   * @returns {string|null}
   */
  static validateField(fieldName, value) {
    let error = null;
    const emailRegex = /^[^\s@]+@[^\s@]+\.[^\s@]+$/;

    switch (fieldName) {
      case "name":
        if (!value) error = "Nome é obrigatório.";
        else if (/\d/.test(value)) error = "Nome não pode conter números.";
        break;
      case "cpf":
        const cpfDigits = value ? value.replace(/\D/g, "") : "";
        if (!cpfDigits) error = "CPF é obrigatório.";
        else if (cpfDigits.length !== 11) error = "CPF deve ter 11 dígitos.";
        break;
      case "registerNumber":
        if (!value) error = "Nº de Registro é obrigatório.";
        else if (value.length !== 7)
          error = "Nº de Registro deve ter 7 dígitos.";
        break;
      case "email":
        if (!value) error = "E-mail é obrigatório.";
        else if (!emailRegex.test(value)) error = "Formato de e-mail inválido.";
        break;
      case "roleId":
        if (!value) error = "ID da Função é obrigatório.";
        break;
    }
    return error;
  }

  /**
   * Valida todos os campos obrigatórios do formulário.
   * @param {object} formData
   * @returns {object}
   */
  static validateAll(formData) {
    const errors = {};
    const { name, cpf, email, registerNumber, roleId } = formData;

    errors.name = User.validateField("name", name);
    errors.cpf = User.validateField("cpf", cpf);
    errors.email = User.validateField("email", email);
    errors.registerNumber = User.validateField(
      "registerNumber",
      registerNumber,
    );
    errors.roleId = User.validateField("roleId", roleId);

    return errors;
  }

  /**
   * Compara o usuário original com os dados do formulário e retorna
   * um objeto contendo apenas os campos que mudaram.
   * @param {object} originalUser
   * @param {object} newFormData
   * @returns {object}
   */
  static getChangedFields(originalUser, newFormData) {
    const updatedUserData = {};

    const { name, email, cpf, registerNumber, roleId, enterpriseId, photoUri } =
      newFormData;

    const cpfDigits = cpf ? cpf.replace(/\D/g, "") : "";
    const roleIdInt = parseInt(roleId, 10);
    const enterpriseIdInt = enterpriseId ? parseInt(enterpriseId, 10) : null;

    if (name !== originalUser.name) {
      updatedUserData.name = name;
    }
    if (email !== originalUser.email) {
      updatedUserData.email = email;
    }
    if (cpfDigits !== originalUser.cpf) {
      updatedUserData.cpf = cpfDigits;
    }
    if (registerNumber !== originalUser.register_number?.toString()) {
      updatedUserData.register_number = registerNumber;
    }
    if (roleIdInt !== originalUser.role_id) {
      updatedUserData.role_id = roleIdInt;
    }
    if (enterpriseIdInt !== originalUser.enterprise_id) {
      updatedUserData.enterprise_id = enterpriseIdInt;
    }
    if (photoUri !== originalUser.photo_url) {
      updatedUserData.photo_url = photoUri || "";
    }

    return updatedUserData;
  }

  /**
   * Retorna um objeto simples para ser enviado como JSON (para CRIALÇÂO)
   */
  toJSON() {
    return {
      name: this.name,
      cpf: this.cpf,
      email: this.email,
      register_number: this.registerNumber,
      password: this.passWord,
      role_id: this.roleId,
      enterprise_id: this.enterpriseId,
    };
  }
}

export default User;
