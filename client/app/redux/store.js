import {createStore, applyMiddleware} from "redux";
import thunk from "redux-thunk";
import createSagaMiddleware from "redux-saga";

import rootReducer from "./root-reducer";
import rootSaga from "./root-saga";
import {logoutCompleted} from "./login/actions";
import {forwardTo} from "./utils";

const errUnauthorized = "Error: Unauthorized";

const authInterceptor = ({dispatch}) => (next) => (action) => {
    if (action.error !== undefined && action.error.toString() === errUnauthorized) {
        dispatch(logoutCompleted());
        forwardTo("/login");
    } else {
        next(action);
    }
};

const sagaMiddleware = createSagaMiddleware();
const middlewares = [authInterceptor, thunk, sagaMiddleware];

const bindMiddleware = middleware => {
    if (process.env.NODE_ENV !== "production") {
        const {composeWithDevTools} = require("redux-devtools-extension");
        return composeWithDevTools(applyMiddleware(...middleware));
    }
    return applyMiddleware(...middleware);
};

const store = createStore(
    rootReducer,
    bindMiddleware(middlewares),
);
sagaMiddleware.run(rootSaga);

export {store};
