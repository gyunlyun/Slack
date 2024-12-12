// Cynhyrchwyd y ffeil hon yn awtomatig. PEIDIWCH Â MODIWL
// This file is automatically generated. DO NOT EDIT
import {structs} from '../models';
import {mongo} from '../models';
import {context} from '../models';

export function AddConnection(arg1:structs.DatabaseConnection):Promise<boolean>;

export function AddFingerscanResult(arg1:string,arg2:structs.InfoResult):Promise<boolean>;

export function AddPocscanResult(arg1:string,arg2:structs.VulnerabilityInfo):Promise<boolean>;

export function AddScanTask(arg1:string,arg2:string,arg3:string,arg4:number,arg5:number):Promise<boolean>;

export function ConnectDatabase(arg1:structs.DatabaseConnection):Promise<boolean>;

export function ConnectMongodb(arg1:string,arg2:string,arg3:string):Promise<mongo.Client>;

export function CreateTable():Promise<boolean>;

export function DisconnectDatabase(arg1:string):Promise<boolean>;

export function ExecSqlStatement(arg1:string,arg2:Array<any>):Promise<boolean>;

export function ExportWebReportWithHtml(arg1:string,arg2:Array<string>):Promise<boolean>;

export function ExportWebReportWithJson(arg1:string,arg2:Array<structs.TaskResult>):Promise<boolean>;

export function FetchDatabaseInfoFromOracle():Promise<{[key: string]: Array<string>}>;

export function FetchDatabaseInfoFromPostgres():Promise<{[key: string]: Array<string>}>;

export function FetchDatabaseinfoFromMongodb():Promise<{[key: string]: Array<string>}>;

export function FetchDatabaseinfoFromMysql():Promise<{[key: string]: Array<string>}>;

export function FetchDatabaseinfoFromSqlServer():Promise<{[key: string]: Array<string>}>;

export function FetchTableInfoFromMysql(arg1:string,arg2:string):Promise<structs.RowData>;

export function FetchTableInfoFromOracle(arg1:string,arg2:string):Promise<structs.RowData>;

export function FetchTableInfoFromPostgres(arg1:string,arg2:string):Promise<structs.RowData>;

export function FetchTableInfoFromSqlServer(arg1:string,arg2:string):Promise<structs.RowData>;

export function GetAllDatabaseConnections():Promise<Array<structs.DatabaseConnection>>;

export function GetAllPathsAndTimes():Promise<Array<structs.PathTimes>>;

export function InsertFavGrammarFiled(arg1:string,arg2:string,arg3:string):Promise<boolean>;

export function ReadWebReportWithJson(arg1:string):Promise<structs.WebReport>;

export function RemoveConnection(arg1:string):Promise<boolean>;

export function RemoveFavGrammarFiled(arg1:string,arg2:string,arg3:string):Promise<boolean>;

export function RemovePocscanResult(arg1:string,arg2:string,arg3:string):Promise<boolean>;

export function RemoveScanTask(arg1:string):Promise<boolean>;

export function RenameScanTask(arg1:string,arg2:string):Promise<boolean>;

export function RetrieveAllScanTasks():Promise<Array<structs.TaskResult>>;

export function RetrieveFingerscanResults(arg1:string):Promise<Array<structs.InfoResult>>;

export function RetrievePocscanResults(arg1:string):Promise<Array<structs.VulnerabilityInfo>>;

export function SearchAgentPool():Promise<Array<string>>;

export function SelectAllSyntax(arg1:string):Promise<Array<structs.SpaceEngineSyntax>>;

export function Startup(arg1:context.Context):Promise<void>;

export function UpdateConnection(arg1:structs.DatabaseConnection):Promise<boolean>;

export function UpdateOrInsertPath(arg1:string):Promise<boolean>;

export function UpdateScanTaskWithResults(arg1:string,arg2:number,arg3:number):Promise<boolean>;
