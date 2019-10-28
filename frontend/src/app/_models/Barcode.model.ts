import {ProductModel} from '.';

export interface BarcodeModel {
  id?: string;
  ean: string;
  product: ProductModel;
  createdAt?: string;
  updatedAt?: string;
}
