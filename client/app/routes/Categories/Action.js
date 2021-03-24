import {
    Button,
    Card,
    CardBody,
    CardTitle,
    Col,
    Form,
    FormGroup,
    Input,
    Row
} from '../../components';
import React from 'react';
import PropTypes from 'prop-types';

export class Action extends React.Component {
    constructor(props) {
        super(props);
    }

    render() {
        const isEdit = this.props.chosenCategory.id > 0;
        const {
            chosenCategory,
            handleResetChosenCategory,
            handleChange,
            handleSubmitCategory,
        } = this.props;

        return (
            <Card className="mb-3">
                <CardBody>
                    <CardTitle tag="h6" className="mb-4">
                        {isEdit ? (
                            <Row>
                                <Col sm={12} lg={7}>
                                    Edit Category
                                </Col>
                                <Col sm={12} lg={3}>
                                    <Button color="primary" onClick={handleResetChosenCategory}>Reset Form</Button>
                                </Col>
                            </Row>
                        ) : (`Create Category`)}
                    </CardTitle>
                    <Form>
                        <FormGroup>
                            <Input
                                type="text"
                                value={chosenCategory.name}
                                onChange={(e) => handleChange(e)}
                                name="name"
                                placeholder="Category name"
                            />
                        </FormGroup>
                        <Button onClick={(e) => handleSubmitCategory(e)} color="primary">
                            {
                                isEdit ? (`Edit`) : (`Create`)
                            }
                        </Button>
                    </Form>
                </CardBody>
            </Card>
        );
    }
}

Action.propTypes = {
    chosenCategory: PropTypes.object.isRequired,
    handleChange: PropTypes.func.isRequired,
    handleResetChosenCategory: PropTypes.func.isRequired,
    handleSubmitCategory: PropTypes.func.isRequired,
};
