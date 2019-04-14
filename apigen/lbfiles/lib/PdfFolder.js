async getCount(PINProjectID, ParentFolderID) { 
		
	if (PINProjectID) {
		
		try { 
		  
			let command = [];
		
			command.push({key: APIVariable.PINProjectID, value: PINProjectID});
		
			command.push({key: APIVariable.ParentFolderID, value: ParentFolderID});
		
		const response = await dataAPI.getDataAPIProjectDB(QueryID.usp_AM_PDFReportFolder_q_case1S_1, PINProjectID, command);
		} catch (error) {
			return error;
		}
		
	} else {
		return util.buildResponse({error: true,
		message: 'getCount - required fields are undefined'});
	}
	
};