import * as React from "react";

export class Hello extends React.Component {
    render() {
        return <h1>{this.props.message}</h1>;
    }
}
