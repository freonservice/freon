import {
    LOGIN_REQUEST,
    LOGOUT_REQUEST,
    LOGOUT_COMPLETED,
} from "./constants";

export const loginRequest = (email, password) => ({
    type: LOGIN_REQUEST,
    email: email,
    password: password,
});

export const logoutRequest = () => {
    return {
        type: LOGOUT_REQUEST,
    };
};

export const logoutCompleted = () => {
    return {
        type: LOGOUT_COMPLETED,
    };
};
