import React from 'react';
import faker from 'faker/locale/en_US';

const Profile = () => {
    return (
        <React.Fragment>
            <div className="mb-4 text-center">
                <a className="h6 text-decoration-none" href="#">
                    { faker.name.firstName() } { faker.name.lastName() }
                </a>
                <div className="text-center mt-2">
                    { faker.name.jobTitle() }
                </div>
                <div className="text-center">
                    <i className="fa fa-map-marker mr-1"></i>
                    { faker.address.city() }
                </div>
            </div>
        </React.Fragment>
    )
}

export { Profile };
