import React, { createContext, useState, useContext, useCallback } from "react";
import ModalErrors from "../components/ModalErrors.jsx";

const ErrorContext = createContext();

export function ErrorProvider({ children }) {
  const [errorState, setErrorState] = useState({
    visible: false,
    message: "",
    onRetry: null,
  });

  const showErrorModal = useCallback((message, onRetryCallback) => {
    setErrorState({
      visible: true,
      message: message || "Ocorreu um erro inesperado.",
      onRetry: onRetryCallback,
    });
  }, []);

  const hideErrorModal = useCallback(() => {
    setErrorState((prev) => ({ ...prev, visible: false }));
  }, []);

  const value = {
    showErrorModal,
    hideErrorModal,
  };

  return (
    <ErrorContext.Provider value={value}>
      {children}
      <ModalErrors
        visible={errorState.visible}
        message={errorState.message}
        onClose={hideErrorModal}
        onRetry={errorState.onRetry ? () => {
          hideErrorModal();
          errorState.onRetry();
        } : null}
      />
    </ErrorContext.Provider>
  );
}

export function useError() {
  const context = useContext(ErrorContext);
  if (!context) {
    throw new Error("useError deve ser usado dentro de um ErrorProvider");
  }
  return context;
}
