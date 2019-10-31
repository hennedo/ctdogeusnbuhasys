import { Injectable } from '@angular/core';
import {HttpClient} from '@angular/common/http';
import {Observable} from 'rxjs';
import {ProductModel} from '../_models';
import {environment} from '../../environments/environment';
import {CurrentUserService} from './current-user.service';

@Injectable({
  providedIn: 'root'
})
export class ProductService {

  constructor(private http: HttpClient, private currentUser: CurrentUserService) { }

  getOne(id: string): Observable<ProductModel> {
    return this.http.get<ProductModel>(environment.apiBaseUrl + '/product/' + id);
  }
  getAll(): Observable<ProductModel[]> {
    return this.http.get<ProductModel[]>(environment.apiBaseUrl + '/product');
  }
  update(id: string, product: ProductModel): Observable<ProductModel> {
    return this.http.put<ProductModel>(environment.apiBaseUrl + '/product/' + id, product, this.currentUser.getHeader());
  }
  create(product: ProductModel): Observable<ProductModel> {
    return this.http.post<ProductModel>(environment.apiBaseUrl + '/product', product, this.currentUser.getHeader());
  }
}
