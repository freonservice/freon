import React from 'react';
import {
    ButtonGroup,
    Button, Card, CardBody, Table//, ModalHeader, ModalBody, ModalFooter
} from './../../components';
import PropTypes from 'prop-types';


export class CategoryTable extends React.Component {
    render() {
        const {listCategories, handleChosenCategory, handleDeleteCategory} = this.props;
        return (
            <Card className="mb-3">
                <CardBody>
                    <Table className="mb-0" bordered responsive>
                        <thead>
                        <tr>
                            <th>Name</th>
                            <th className="text-right">
                                Actions
                            </th>
                        </tr>
                        </thead>
                        <tbody>
                        <React.Fragment>
                            {
                                listCategories.length > 0 && listCategories.map((product, index) => (
                                        <tr key={index}>
                                            <td className="align-middle">
                                            <span className="text-inverse">
                                                {product.name}
                                            </span>
                                            </td>
                                            <td className="text-right">
                                                <ButtonGroup>
                                                    <Button
                                                        color="link"
                                                        onClick={() => handleChosenCategory(index)}
                                                        className="text-decoration-none"
                                                    >
                                                        <i className="fa fa-edit"/>
                                                    </Button>
                                                    <Button
                                                        color="link"
                                                        onClick={() => handleDeleteCategory(product.id)}
                                                        className="text-decoration-none"
                                                    >
                                                        <i className="fa fa-close"/>
                                                    </Button>
                                                </ButtonGroup>
                                            </td>
                                        </tr>
                                    )
                                )
                            }
                        </React.Fragment>
                        </tbody>
                    </Table>
                </CardBody>
            </Card>
        );
    }
}

CategoryTable.propTypes = {
    listCategories: PropTypes.array,
    handleChosenCategory: PropTypes.func.isRequired,
    handleDeleteCategory: PropTypes.func.isRequired,
};

CategoryTable.defaultProps = {
    listCategories: []
};
