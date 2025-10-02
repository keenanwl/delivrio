import {Component, OnInit} from '@angular/core';
import {Observable} from "rxjs";
import {AppModel, AppState} from "../app.ngxs";
import {Store} from "@ngxs/store";
import {NestedTreeControl} from "@angular/cdk/tree";
import {MatTreeNestedDataSource} from "@angular/material/tree";
import {Paths} from "../app-routing.module";

interface FoodNode {
	name: string;
	uri: string;
	limitedDisabled: boolean;
	children?: FoodNode[];
}

const TREE_DATA: FoodNode[] = [
	{
		name: 'Users',
		uri: 'users/list',
		limitedDisabled: false,
	},
	{
		name: 'Connections',
		uri: 'connections',
		limitedDisabled: false,
	},
	{
		name: 'Carriers',
		uri: 'carriers',
		limitedDisabled: false,
	},
	{
		name: 'Delivery Options',
		uri: 'delivery-options',
		limitedDisabled: false,
	},
	{
		name: 'Email templates',
		uri: 'email-templates',
		limitedDisabled: true,
	},
	{
		name: 'Notifications',
		uri: 'notifications',
		limitedDisabled: true,
	},
	{
		name: 'Locations',
		uri: 'locations',
		limitedDisabled: true,
	},
	{
		name: 'A/B (Hypothesis) Testing',
		uri: 'hypothesis-testing',
		limitedDisabled: true,
	},
	{
		name: 'Return portals',
		uri: 'return-portals',
		limitedDisabled: true,
	},
	{
		name: 'API Tokens',
		uri: 'api-tokens',
		limitedDisabled: true,
	},
	{
		name: 'Packaging',
		uri: 'packaging',
		limitedDisabled: true,
	},
	{
		name: 'User groups',
		uri: 'users/groups/list',
		limitedDisabled: true,
	},
	{
		name: 'Workstations',
		uri: 'workstations',
		limitedDisabled: true,
	},
];

@Component({
	selector: 'app-settings',
	templateUrl: './settings.component.html',
	styleUrls: ['./settings.component.scss'],
})
export class SettingsComponent implements OnInit {
	treeControl = new NestedTreeControl<FoodNode>(node => node.children);
	dataSource = new MatTreeNestedDataSource<FoodNode>();

	app$: Observable<AppModel>;

	constructor(private store: Store) {
		this.app$ = store.select(AppState.get);
		this.dataSource.data = TREE_DATA;
	}

	ngOnInit(): void {
	}

	hasChild = (_: number, node: FoodNode) => !!node.children && node.children.length > 0;

}
