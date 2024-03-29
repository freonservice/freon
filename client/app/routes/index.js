import React from 'react';
import {
    Route,
    Switch,
    Redirect
} from 'react-router';

import Dashboard from './Dashboard';
import Categories from './Categories';
import Localizations from './Localizations';
import Identifiers from './Identifiers';
import Translations from './Translations';
import TranslationAction from "./Translations/TranslationAction";
import TranslationFiles from "./TranslationFiles";
import UsersList from "./Users/Users";
import CreateUser from './Users/CreateUser';
import EditUser from "./Users/EditUser";
import Login from './Login';

import Profile from "./Settings/Profile";
import Security from "./Settings/Security";
import Storage from "./Settings/Storage";
import Translation from "./Settings/Translation";

import PrivateRoute from "./privateroute";

export const RoutedContent = () => {
    return (
        <Switch>
            <Redirect from="/" to="/dashboard" exact/>

            <PrivateRoute path="/dashboard" exact component={Dashboard}/>

            <PrivateRoute path='/localizations' exact component={Localizations}/>
            <PrivateRoute path='/categories' exact component={Categories}/>
            <PrivateRoute path='/identifiers' exact component={Identifiers}/>

            <PrivateRoute path='/translations' exact component={Translations}/>
            <PrivateRoute path='/translations/edit/:id' exact component={TranslationAction}/>
            <PrivateRoute path='/translations/create' exact component={TranslationAction}/>

            <PrivateRoute path='/translation-files' exact component={TranslationFiles}/>

            <PrivateRoute path='/users' exact component={UsersList}/>
            <PrivateRoute path='/users/edit/:id' exact component={EditUser}/>
            <PrivateRoute path='/users/create' exact component={CreateUser}/>

            <PrivateRoute path='/settings/profile-edit' exact component={Profile}/>
            <PrivateRoute path='/settings/security-edit' exact component={Security}/>
            <PrivateRoute path='/settings/storage-edit' exact component={Storage}/>
            <PrivateRoute path='/settings/translation-edit' exact component={Translation}/>

            <Route path="/login" exact component={Login}/>

            <Redirect to="/pages/error-404"/>
        </Switch>
    );
};
