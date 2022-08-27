import grpc from '@grpc/grpc-js'
import { ExecutorClient } from './pb/cowait_grpc_pb'
import { Header, TaskInitReq } from './pb/cowait_pb'
import { Timestamp } from 'google-protobuf/google/protobuf/timestamp_pb';


export class Client {
    private task_id: string
    private client: ExecutorClient

    public constructor(endpoint: string, id: string) {
        this.task_id = id
        this.client = new ExecutorClient(endpoint, grpc.credentials.createInsecure())
    }

    private header(): Header {
        let header = new Header()
        header.setId(this.task_id)

        let ts = new Timestamp()
        ts.fromDate(new Date())
        header.setTime(ts)

        return header
    }

    public async init(){
        let req = new TaskInitReq()
        req.setHeader(this.header())
        req.setVersion('cowait-js/1.0')
        return new Promise<void>((resolve, reject) => {
            this.client.taskInit(req, (err, _) => {
                if (err) {
                    reject(err)
                    return
                }
                resolve()
            })
        })
    }
}
