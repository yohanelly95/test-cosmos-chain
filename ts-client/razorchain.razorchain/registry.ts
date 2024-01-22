import { GeneratedType } from "@cosmjs/proto-signing";
import { GenesisState } from "./types/razorchain/razorchain/genesis";
import { QueryParamsRequest } from "./types/razorchain/razorchain/query";
import { MsgUpdateParams } from "./types/razorchain/razorchain/tx";
import { MsgUpdateParamsResponse } from "./types/razorchain/razorchain/tx";
import { Params } from "./types/razorchain/razorchain/params";
import { QueryParamsResponse } from "./types/razorchain/razorchain/query";

const msgTypes: Array<[string, GeneratedType]>  = [
    ["/razorchain.razorchain.GenesisState", GenesisState],
    ["/razorchain.razorchain.QueryParamsRequest", QueryParamsRequest],
    ["/razorchain.razorchain.MsgUpdateParams", MsgUpdateParams],
    ["/razorchain.razorchain.MsgUpdateParamsResponse", MsgUpdateParamsResponse],
    ["/razorchain.razorchain.Params", Params],
    ["/razorchain.razorchain.QueryParamsResponse", QueryParamsResponse],
    
];

export { msgTypes }