import { ComponentFixture, TestBed } from '@angular/core/testing';

import { StoreMapSelectorComponent } from './store-map-selector.component';

describe('StoreMapSelectorComponent', () => {
  let component: StoreMapSelectorComponent;
  let fixture: ComponentFixture<StoreMapSelectorComponent>;

  beforeEach(() => {
    TestBed.configureTestingModule({
      declarations: [StoreMapSelectorComponent]
    });
    fixture = TestBed.createComponent(StoreMapSelectorComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
