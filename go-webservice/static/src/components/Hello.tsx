import * as React from "react";

export interface HelloProps { message: string }

export class Hello extends React.Component<HelloProps, {}> {
    render() {
        return <h1>{this.props.message}</h1>;
    }
}
