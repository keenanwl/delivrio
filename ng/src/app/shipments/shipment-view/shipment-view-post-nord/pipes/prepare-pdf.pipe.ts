import {Pipe, PipeTransform} from '@angular/core';

@Pipe({
	name: 'shipmentPreparePdf',
	standalone: true,
})
export class ShipmentPreparePdfPipe implements PipeTransform {

  transform(value: string): Uint8Array {
	  const raw = atob(value);
	  const uint8Array = new Uint8Array(raw.length);
	  for (let i = 0; i < raw.length; i++) {
		  uint8Array[i] = raw.charCodeAt(i);
	  }
	  return uint8Array;
  }

}
