import {CategoryModel} from '.';

export interface ProductModel {
  id?: string;
  name: string;
  price: number;
  stock: number;
  customAttributes: string;
  category: CategoryModel;
  createdAt?: string;
  updatedAt?: string;
}
