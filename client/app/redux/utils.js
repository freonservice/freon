import history from "../history";

export function forwardTo(location) {
    history.push(location);
}

export const tokenKey = "token";
const defaultBody = {
    "Content-Type": "application/json",
};

export function getHeaders() {
    const token = JSON.parse(localStorage.getItem(tokenKey));
    if (token !== null && token.length > 0) {
        return {
            ...defaultBody,
            "Authorization": token,
        };
    }
    return defaultBody;
}