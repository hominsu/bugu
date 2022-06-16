<div id="top"></div>

<!-- PROJECT SHIELDS -->
<p align="center">
<a href="https://github.com/hominsu/bugu/graphs/contributors"><img src="https://img.shields.io/github/contributors/hominsu/bugu.svg?style=for-the-badge" alt="Contributors"></a>
<a href="https://github.com/hominsu/bugu/network/members"><img src="https://img.shields.io/github/forks/hominsu/bugu.svg?style=for-the-badge" alt="Forks"></a>
<a href="https://github.com/hominsu/bugu/stargazers"><img src="https://img.shields.io/github/stars/hominsu/bugu.svg?style=for-the-badge" alt="Stargazers"></a>
<a href="https://github.com/hominsu/bugu/issues"><img src="https://img.shields.io/github/issues/hominsu/bugu.svg?style=for-the-badge" alt="Issues"></a>
<a href="https://github.com/hominsu/bugu/blob/master/LICENSE"><img src="https://img.shields.io/github/license/hominsu/bugu.svg?style=for-the-badge" alt="License"></a>
<a href="https://github.com/hominsu/bugu/actions/workflows/docker-publish.yml"><img src="https://img.shields.io/github/workflow/status/hominsu/bugu/Docker%20Deploy?style=for-the-badge" alt="Deploy"></a>
</p>


<!-- PROJECT LOGO -->
<br/>
<div align="center">
<!--   <a href="https://github.com/hominsu/bugu">
    <img src="images/logo.png" alt="Logo" width="80" height="80">
  </a> -->

<h3 align="center">bugu</h3>

  <p align="center">
    基于人工智能检测的免杀系统
    <br/>
    <a href="https://hominsu.github.io/bugu/"><strong>Explore the docs » (you are here)</strong></a>
    <br/>
    <br/>
    <a href="https://github.com/hominsu/bugu">View Demo</a>
    ·
    <a href="https://github.com/hominsu/bugu/issues">Report Bug</a>
    ·
    <a href="https://github.com/hominsu/bugu/issues">Request Feature</a>
  </p>
</div>

## Description

基于人工智能检测的免杀系统

## Details

```mermaid
flowchart LR
	admin("admin service") <-.-> user("user service")
	bugu("bugu service") <-.-> user
	bugu <-.-> detect("detect service")
	bugu <-.-> packer("packer service")
	bugu <-.-> confusion("confusion service")
	
	subgraph DB
	redis[("redis")]
	userdb[("user db")]
	kafka[("kafka")]
	end
	
	subgraph File
	oss[("oss")]
	metadatadb[("file meta db")]
	end
	
	bugu <-.file.-> oss
	admin <-.file.-> oss
	
	bugu <-.file metadata.-> metadatadb
	admin <-.file metadata.-> metadatadb
	
	user <-.user info.-> userdb
	user <-.user cache.-> redis
	
	bugu -.delay task.-> kafka
	detect <-.delay task.- kafka
	packer <-.delay task.- kafka
	confusion <-.delay task.- kafka
	
	subgraph Other Infrastructure
	consul("Consul")
	sls("Aliyun Log Service")
	end
	
```
