import {useDispatch} from "react-redux";

function useDispatchPromise() {
    const dispatch = useDispatch();
    return props => {
        return new Promise((resolve, reject) => {
            dispatch({...props, resolve, reject});
        });
    };
}

export default useDispatchPromise;
