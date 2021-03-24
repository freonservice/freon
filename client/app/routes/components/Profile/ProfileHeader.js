import React from 'react';
import { Link } from 'react-router-dom';

import { 
    Badge,
    Media,
} from './../../../components';

const ProfileHeader = (profile) => (
    <React.Fragment>
        <Media className="mb-3">
            <Media body>
                <h5 className="mb-1 mt-0">
                    <Link to="/apps/profile-details">
                        { profile.first_name } { profile.second_name }
                    </Link>
                </h5>
            </Media>
        </Media>
    </React.Fragment>
)

export { ProfileHeader };
