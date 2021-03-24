import {
    LOGIN_REQUEST,
    LOGIN_SUCCESS,
    LOGIN_ERROR,
    LOGOUT_COMPLETED,
    LOGOUT_ERROR,
} from "./constants";

import {tokenKey} from "../utils";

const initialState = {
    isLoading: false,
    successful: false,
    error: "",
};

export default function reducer(state = initialState, action) {
    switch (action.type) {
        case LOGIN_REQUEST:
            return {
                ...state,
                isLoading: true,
            };
        case LOGIN_SUCCESS:
            return {
                error: "",
                isLoading: false,
                successful: true,
            };
        case LOGIN_ERROR:
            return {
                ...state,
                error: action.error.toString(),
                isLoading: false,
                successful: false,
            };
        case LOGOUT_COMPLETED:
            console.log("LOGOUT_COMPLETED");
            localStorage.removeItem(tokenKey);
            return {
                error: "",
                isLoading: false,
                successful: false,
            };
        case LOGOUT_ERROR:
            return {
                ...state,
                error: action.error.toString(),
                isLoading: false,
                successful: false,
            };
        default:
            return state;
    }
}
