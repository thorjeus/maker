// Copyright (C) 2018 Cranky Kernel
//
// This program is free software: you can redistribute it and/or modify
// it under the terms of the GNU Affero General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.
//
// This program is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
// GNU Affero General Public License for more details.
//
// You should have received a copy of the GNU Affero General Public License
// along with this program. If not, see <http://www.gnu.org/licenses/>.

import {Injectable} from "@angular/core";

@Injectable({
    providedIn: "root",
})
export class LoggerService {

    constructor() {
    }

    log(msg) {
        console.log(msg);
    }

    logWithPrefix(prefix, msg) {
        switch (typeof msg) {
            case "string":
                this.log(`${prefix}: ${msg}`);
                break;
            default:
                this.log(`${prefix}: ${JSON.stringify(msg)}`);
                break;
        }
    }

    getLogger(prefix: string): Logger {
        return new Logger(prefix, this);
    }
}

export class Logger {
    constructor(private prefix: string, private logger: LoggerService) {
    }

    log(msg) {
        this.logger.logWithPrefix(this.prefix, msg);
    }
}
