import React from 'react';
import {Link} from 'react-router-dom';

import {
    Sidebar,
    UncontrolledButtonDropdown,
    Avatar,
    AvatarAddOn,
    DropdownToggle,
    DropdownMenu,
    DropdownItem
} from '../../../components';
import {connect} from 'react-redux';
import {profileRequest} from '../../../redux/profile/actions';
import * as PropTypes from 'prop-types';
import {logoutRequest} from '../../../redux/login/actions';

class SidebarProfile extends React.Component {
    constructor(props) {
        super(props);

        this.props.profileRequest();
    }

    render() {
        return (
            <React.Fragment>
                <Sidebar.HideSlim>
                    <Sidebar.Section className="pt-0">
                        <Link to="/" className="d-block">
                            <Sidebar.HideSlim>
                                <Avatar.Image
                                    size="lg"
                                    src="https://cdn.pixabay.com/photo/2016/11/18/23/38/child-1837375_1280.png"
                                />
                            </Sidebar.HideSlim>
                        </Link>
                        <UncontrolledButtonDropdown>
                            <DropdownToggle color="link" className="pl-0 pb-0 btn-profile sidebar__link">
                                {this.props.profile.first_name} {this.props.profile.second_name}
                                <i className="fa fa-angle-down ml-2"/>
                            </DropdownToggle>
                            <DropdownMenu persist>
                                <DropdownItem header>
                                    {this.props.profile.first_name} {this.props.profile.second_name}
                                </DropdownItem>
                                <DropdownItem divider/>
                                <DropdownItem tag={Link} to="/settings/profile-edit">
                                    My Profile
                                </DropdownItem>
                                <DropdownItem tag={Link} to="/settings/security-edit">
                                    Settings
                                </DropdownItem>
                                <DropdownItem divider/>
                                <DropdownItem onClick={this.props.logoutRequest}>
                                    <i className="fa fa-fw fa-sign-out mr-2"/>
                                    Sign Out
                                </DropdownItem>
                            </DropdownMenu>
                        </UncontrolledButtonDropdown>
                        <div className="small sidebar__link--muted">
                            Administrator
                        </div>
                    </Sidebar.Section>
                </Sidebar.HideSlim>
                <Sidebar.ShowSlim>
                    <Sidebar.Section>
                        <Avatar.Image
                            size="sm"
                            src="https://cdn.pixabay.com/photo/2016/11/18/23/38/child-1837375_1280.png"
                            addOns={[
                                <AvatarAddOn.Icon
                                    className="fa fa-circle"
                                    color="white"
                                    key="avatar-icon-bg"
                                />,
                                <AvatarAddOn.Icon
                                    className="fa fa-circle"
                                    color="success"
                                    key="avatar-icon-fg"
                                />
                            ]}
                        />
                    </Sidebar.Section>
                </Sidebar.ShowSlim>
            </React.Fragment>
        );
    }
}

SidebarProfile.propTypes = {
    profileRequest: PropTypes.func.isRequired,
    logoutRequest: PropTypes.func.isRequired,
    profile: PropTypes.shape({
        first_name: PropTypes.string,
        second_name: PropTypes.string,
    }).isRequired,
};

const mapStateToProps = (state) => ({
    profile: state.profile.profile
});

export default connect(mapStateToProps, {profileRequest, logoutRequest})(SidebarProfile);