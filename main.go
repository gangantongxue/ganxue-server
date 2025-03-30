/**
 * Copyright 2025-2025 gangantongxue. All Rights Reserved.

 * Licensed under the GPL License, Version 3.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at

 *      https://www.gnu.org/licenses/gpl-3.0.html

 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

package main

import (
	"ganxue-server/initialize"
	"ganxue-server/router"
	_ "net/http/pprof"
)

func main() {
	// 初始化
	initialize.InitAll()
	defer initialize.CloseAll()

	// 初始化路由
	h := initialize.Routers()
	router.Init(h)

	// 启动服务
	h.Spin()
}
