import React from 'react';

import {Container} from './../../components';
import {HeaderMain} from "../components/HeaderMain";
import faker from "faker/locale/en_US";
import _ from "lodash";
import {GenerationTable} from "./Table";
import {ToastContainer} from "react-toastify";

const generateRow = (index) => ({
    id: index,
    local_name: faker.commerce.productName(),
    lang_name: faker.commerce.productName(),
    lcid: faker.commerce.productName(),
});

class GenerationFiles extends React.Component {
    constructor(props) {
        super(props);

        this.state = {
            listLocalization: _.times(10, generateRow),
        };
    }

    render() {
        return (
            <Container>
                <HeaderMain
                    title="Generation Files"
                    className="mb-5 mt-4"
                />
                <GenerationTable
                    listLocalization={this.state.listLocalization}
                />
                <ToastContainer
                    position="top-right"
                    autoClose={3000}
                    draggable={false}
                    hideProgressBar={true}
                    limit={3}
                />
            </Container>
        )
    }
}

export default GenerationFiles;