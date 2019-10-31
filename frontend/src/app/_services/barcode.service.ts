import { Injectable } from '@angular/core';
import {HttpClient} from '@angular/common/http';
import {CurrentUserService} from './current-user.service';
import {Observable} from 'rxjs';
import {BarcodeModel} from '../_models';
import {environment} from '../../environments/environment';

@Injectable({
  providedIn: 'root'
})
export class BarcodeService {

  constructor(private http: HttpClient, private currentUser: CurrentUserService) { }

  getOne(id: string): Observable<BarcodeModel> {
    return this.http.get<BarcodeModel>(environment.apiBaseUrl + '/barcode/' + id);
  }
  getAll(): Observable<BarcodeModel[]> {
    return this.http.get<BarcodeModel[]>(environment.apiBaseUrl + '/barcode');
  }
  update(id: string, barcode: BarcodeModel): Observable<BarcodeModel> {
    return this.http.put<BarcodeModel>(environment.apiBaseUrl + '/barcode/' + id, barcode, this.currentUser.getHeader());
  }
  create(barcode: BarcodeModel): Observable<BarcodeModel> {
    return this.http.post<BarcodeModel>(environment.apiBaseUrl + '/barcode', barcode, this.currentUser.getHeader());
  }
}
