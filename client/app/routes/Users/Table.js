import React from 'react';
import BootstrapTable from 'react-bootstrap-table-next';
import ToolkitProvider from 'react-bootstrap-table2-toolkit';
import moment from 'moment';

import {
    Button,
    ButtonGroup,
} from './../../components';
import * as PropTypes from 'prop-types';
import CustomSearch from '../../components/CustomSearch';

const sortCaret = (order) => {
    if (order) {
        return <i className={`fa fa-fw text-muted fa-sort-${order}`}/>;
    } else {
        return <i className="fa fa-fw fa-sort text-muted"/>;
    }
};

export class Table extends React.Component {
    constructor(props) {
        super(props);

        this.state = {
            users: this.props.listUsers
        };
    }

    handleCreateUser = () => {
        this.props.history.push({
            pathname: `/users/create`,
        });
    };

    handleEditUser = (id) => {
        const user = this.state.users.find(o => o.id === id);
        this.props.history.push({
            pathname: `/users/edit/${id}`,
            state: user
        });
    };

    createColumnDefinitions() {
        return [
            {
                dataField: 'id',
                text: 'ID',
                sort: true,
                sortCaret
            }, {
                dataField: 'first_name',
                text: 'First Name',
                sort: true,
                sortCaret
            }, {
                dataField: 'second_name',
                text: 'Last Name',
                sort: true,
                sortCaret
            }, {
                dataField: 'email',
                text: 'Email',
                sort: true,
                sortCaret
            }, {
                dataField: 'role',
                text: 'Role',
                sort: true,
                sortCaret,
                formatter: (cell, row) => {
                    return row.role.charAt(0).toUpperCase() + row.role.slice(1); // Uppercase first symbol
                },
            }, {
                dataField: 'status',
                text: 'Status',
                sort: true,
                sortCaret,
                formatter: (cell, row) => {
                    return row.status.charAt(0).toUpperCase() + row.status.slice(1); // Uppercase first symbol
                },
            }, {
                dataField: 'created_at',
                text: 'Created',
                sort: true,
                sortCaret,
                formatter: (cell, row) => {
                    return (moment.unix(row.created_at).format('DD-MMM-YYYY'));
                },
            }, {
                dataField: 'uuid_id',
                text: 'Action',
                sort: false,
                formatter: (cell, row) => {
                    return (
                        <ButtonGroup>
                            <Button onClick={this.handleEditUser.bind(this, row.id)} color="primary">
                                Edit
                            </Button>
                        </ButtonGroup>
                    );
                },
            }
        ];
    }

    render() {
        const columnDefs = this.createColumnDefinitions();

        return (
            <ToolkitProvider
                keyField="id"
                data={this.state.users}
                columns={columnDefs}
                search
            >
                {
                    props => (
                        <React.Fragment>
                            <div className="d-flex justify-content-end align-items-center mb-2">
                                <div className="d-flex ml-auto">
                                    <CustomSearch
                                        className="mr-2"
                                        {...props.searchProps}
                                    />
                                    <ButtonGroup>
                                        <Button
                                            size="sm"
                                            outline
                                            onClick={this.handleCreateUser}
                                        >
                                            Create <i className="fa fa-fw fa-plus"/>
                                        </Button>
                                    </ButtonGroup>
                                </div>
                            </div>
                            <BootstrapTable
                                classes="table-responsive-lg"
                                bordered={true}
                                responsive
                                hover
                                {...props.baseProps}
                            />
                        </React.Fragment>
                    )
                }
            </ToolkitProvider>
        );
    }
}

Table.propTypes = {
    listUsers: PropTypes.array,
    history: PropTypes.shape({
        push: PropTypes.func.isRequired,
    }).isRequired,
};

Table.defaultProps = {
    listUsers: []
};
