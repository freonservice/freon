import React from 'react';
import {hot} from 'react-hot-loader'
import {Router} from 'react-router-dom';
import {Provider} from 'react-redux';
import AppLayout from './../../layout/default';
import {RoutedContent} from '../../routes';
import history from '../../history';
import {store} from '../../redux/store';

const AppClient = () => {
    return (
        <Provider store={store}>
            <Router history={history}>
                <AppLayout>
                    <RoutedContent/>
                </AppLayout>
            </Router>
        </Provider>
    );
}

export default hot(module)(AppClient);