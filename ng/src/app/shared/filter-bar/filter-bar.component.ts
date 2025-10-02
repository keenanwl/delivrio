import {Component, EventEmitter, Input, Output, ViewChild} from '@angular/core';
import {Subscription, timer} from "rxjs";
import {Store} from "@ngxs/store";
import {MatDialog} from "@angular/material/dialog";
import {FormControl} from "@angular/forms";
import {MatAutocompleteTrigger} from "@angular/material/autocomplete";
type dropDownItem = {name: string; id: string; icon?: string};
type dropDownList = dropDownItem[];

type filterCategory = dropDownItem;
export type filterCategoryList = filterCategory[];
type selectedOption = {filterName: string; filterID: string; optionID: string; optionName: string};
export type selectedOptionList = selectedOption[];

type searchInput = dropDownItem;
export type searchInputList = searchInput[];

export type activeFilter = {filterID: string; lookup: string};

@Component({
	selector: 'app-filter-bar',
	templateUrl: './filter-bar.component.html',
	styleUrls: ['./filter-bar.component.scss']
})
export class FilterBarComponent {
	@Input()
	get searchOptions(): searchInputList {
		return this._searchOptions;
	}

	set searchOptions(value: searchInputList) {
		this._searchOptions = value;
		if (!!this.active) {
			this.subscriptions.push(timer(10).subscribe(() => this.autocomplete?.openPanel()));
		}
		timer(250).subscribe(() => this.loading = false);
	}
	_searchOptions: searchInputList = [];

	@ViewChild("autoInput", {read: MatAutocompleteTrigger}) autocomplete: MatAutocompleteTrigger | undefined;

	@Input() filterCategories: filterCategoryList = [];

	@Output() activeFilter: EventEmitter<activeFilter> = new EventEmitter();
	@Output() selectedFilters: EventEmitter<selectedOptionList> = new EventEmitter();

	active: dropDownItem | null = null;
	allSelected: selectedOptionList = [];
	loading = false;

	filterControl = new FormControl('');

	subscriptions: Subscription[] = [];

	constructor(
		private store: Store,
		private dialog: MatDialog,
	) {
	}

	// Filter categories and options use the same dropdown
	// so we require a fn to handle which one to display
	getDropdownData(filterCategories: filterCategoryList, searchOptions: searchInputList): dropDownList {
		if (!this.active) {
			return filterCategories;
		}
		return searchOptions;
	}

	dropDownSelected(i: dropDownItem) {
		this.filterControl.setValue(``)
		this.allSelected = this.allSelected.filter((s) => !!this.active && s.filterID !== this.active?.id);
		if (!!this.active) {
			this.allSelected.push({
				filterID: this.active.id,
				filterName: this.active.name,
				optionID: i.id,
				optionName: i.name,
			})
			this.active = null;
			this.selectedFilters.emit([...this.allSelected]);
			return
		}

		this.loading = true;
		this.allSelected.push({
			filterID: i.id,
			filterName: i.name,
			optionID: "",
			optionName: "",
		});
		this.active = i;
		this.activeFilter.emit({filterID: i.id, lookup: ''});
	}

	inputChange(val: string) {
		if (!!this.active) {
			this.loading = true;
			this.activeFilter.emit({filterID: this.active.id, lookup: val});
		}
	}

	chipClickSelect(item: selectedOption) {
		this.allSelected = this.allSelected.filter((s) => !!this.active && s.filterID !== this.active?.id);
		this.active = null;
		this.dropDownSelected({id: item.filterID, name: item.filterName});
	}

	removeChip(filterID: string) {
		this.active = null;
		this.allSelected = this.allSelected.filter((f) => f.filterID !== filterID);
		this.selectedFilters.emit([...this.allSelected]);
	}

	removeAllChips() {
		this.active = null;
		this.allSelected = [];
		this.selectedFilters.emit([...this.allSelected]);
		this.filterControl.reset();
	}

}
