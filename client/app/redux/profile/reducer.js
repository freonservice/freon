import {
    PROFILE_REQUEST,
    PROFILE_ERROR,
    PROFILE_SUCCESS,
} from "./constants";

const initialState = {
    isLoading: false,
    successful: false,
    error: "",
    profile: {},
};

export default function reducer(state = initialState, action) {
    switch (action.type) {
        case PROFILE_REQUEST:
            return {
                ...state,
                isLoading: true,
            };
        case PROFILE_SUCCESS:
            return {
                error: "",
                isLoading: false,
                successful: true,
                profile: action.payload,
            };
        case PROFILE_ERROR:
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
