async getCount(PINProjectID, ParentFolderID) { 
		
	if (PINProjectID) {
		
		try { 
		  
			let command = [];
		
			command.push({key: APIVariable.PINProjectID, value: PINProjectID});
		
			command.push({key: APIVariable.ParentFolderID, value: ParentFolderID});
		
			const response = await dataAPI.getDataAPIProjectDB(QueryID.usp_AM_PDFReportFolder_q_case1S_1, PINProjectID, command);
			if(response && response.result.length > 0) {
				return response;
			} else {
				logger.log('error','ReportFolder.getCount- Error - Unable to process');
			}
			
		
		} catch (error) {
			logger.log('error',`ReportFolder.getCount- Error - ${error.message}`);	
			return Promise.resolve(error);
		}
	} else {
		return util.buildResponse({error: true,
		message: 'getCount - required fields are undefined'});
	}
	
};