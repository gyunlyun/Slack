// Cynhyrchwyd y ffeil hon yn awtomatig. PEIDIWCH Â MODIWL
// This file is automatically generated. DO NOT EDIT
import {main} from '../models';
import {clients} from '../models';
import {space} from '../models';
import {portscan} from '../models';
import {poc} from '../models';

export function AssetHunter(arg1:number,arg2:string,arg3:string):Promise<main.HunterSearch>;

export function AssetSubcompany(arg1:string,arg2:number):Promise<main.SubcompanyInfo>;

export function AssetWechat(arg1:string):Promise<main.WechatInfo>;

export function CheckCdn(arg1:string,arg2:string,arg3:string):Promise<string>;

export function CheckFileStat(arg1:string):Promise<boolean>;

export function CheckTarget(arg1:string,arg2:clients.Proxy):Promise<main.AliveTarget>;

export function DefaultOpenURL(arg1:string):Promise<void>;

export function DomainInfo(arg1:string):Promise<string>;

export function ExtractIP(arg1:string):Promise<string>;

export function FingerScan(arg1:string,arg2:Array<number>,arg3:clients.Proxy):Promise<main.InfoResult>;

export function FofaSearch(arg1:string,arg2:string,arg3:string,arg4:string,arg5:string,arg6:boolean,arg7:boolean):Promise<space.FofaSearchResult>;

export function FofaTips(arg1:string):Promise<space.TipsResult>;

export function Fscan2Txt(arg1:string):Promise<string>;

export function GetFileContent(arg1:string):Promise<string>;

export function GetFingerPoc(arg1:Array<string>):Promise<Array<string>>;

export function GoSimpleFetch(arg1:string):Promise<main.Response>;

export function HostAlive(arg1:Array<string>,arg2:boolean):Promise<Array<string>>;

export function IPParse(arg1:string):Promise<Array<string>>;

export function IconHash(arg1:string):Promise<string>;

export function InitDict(arg1:Array<string>):Promise<Array<string>>;

export function InitIPResolved():Promise<void>;

export function InitRule():Promise<Array<number>>;

export function InitTycHeader(arg1:string):Promise<void>;

export function LocalWalkFiles(arg1:string):Promise<Array<string>>;

export function Minimise():Promise<void>;

export function PathRequest(arg1:string,arg2:string,arg3:number,arg4:string):Promise<main.PathData>;

export function PocNums(arg1:string,arg2:string):Promise<number>;

export function PortBrute(arg1:string,arg2:Array<string>,arg3:Array<string>):Promise<portscan.Burte>;

export function PortCheck(arg1:string,arg2:number,arg3:number):Promise<main.PortResult>;

export function PortParse(arg1:string):Promise<Array<number>>;

export function Quit():Promise<void>;

export function ReadPocDetail(arg1:string):Promise<poc.VulnerabilityDetails>;

export function Restart():Promise<void>;

export function Subdomain(arg1:string,arg2:string,arg3:string,arg4:number):Promise<Array<string>>;

export function System(arg1:string,arg2:number):Promise<Array<any>>;

export function TestTarget(arg1:string):Promise<boolean>;

export function ThinkDict(arg1:string,arg2:string,arg3:string,arg4:string,arg5:string):Promise<Array<string>>;

export function ToggleMaximise():Promise<void>;

export function Transcoding(arg1:Array<string>,arg2:string,arg3:number):Promise<string>;

export function UpdateClinetFile(arg1:string):Promise<string>;

export function UpdatePocFile(arg1:string):Promise<string>;

export function Webscan(arg1:string,arg2:string,arg3:string,arg4:Array<string>,arg5:clients.Proxy):Promise<any>;

export function WechatAppid(arg1:string,arg2:string):Promise<string>;
