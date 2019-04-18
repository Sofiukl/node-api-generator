async addComment(input) { 
		
	if (input && input.PINProjectID) {
		
		try { 
			const ReportCommentID = await dataAPI.executeSequenceGenerator(TableID.AM_PDFReportComment,1);
			 const currentUser = util.getLoggedInUser();
		
			input.CreateBy = currentUser.PWUserID;
		  
			let command = [];
		
			command.push({key: APIVariable.PDFReportID, value: input.PDFReportID});
		
			command.push({key: APIVariable.Comment, value: input.Comment});
		
			command.push({key: APIVariable.ReportActivity, value: input.ReportActivity});
		
			command.push({key: APIVariable.CreateBy, value: input.CreateBy});
		
			command.push({key: APIVariable.ReportCommentID, value: ReportCommentID.result});
		
			const response = await dataAPI.getDataAPIProjectDB(QueryID.usp_AM_PDFReportComment_i, PINProjectID, command);
			if(response && response.result[0].affectedRows != 1) {
				return Promise.resolve(result);
			}
			await this.updateReportAuditFields(input.PINProjectID, input);
			
			result.ReportCommentID = ReportCommentID.result;
			
			result.ReportActivity = ReportActivity;
			
			return Promise.resolve({
				result: [{
					error: false,
					ReportCommentID: ReportCommentID.result
				}]
			});
		
		} catch (error) {
			logger.log('error',`ReportComment.addComment- Error - ${error.message}`);	
			return Promise.resolve(error);
		}
	} else {
		return util.buildResponse({error: true,
		message: 'addComment - required fields are undefined'});
	}
	
};